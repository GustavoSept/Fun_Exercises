import java.text.NumberFormat;

public class Main {
    public static void main(String[] args) {
        // USER DEFINED
        long initial_Investment = 40000;
        float salary_productiveYears = 7500f;
        float salary_retiredYears = 5000f;
        float monthlySalaryShare = 0.15f; // How much the user invests, as a % of the salary_productiveYears (For most people, should be between 10% to 30%)
        float yield_investment = 0.10f; // Yield Rate during investment period (realistic values for most people are 6~10% (after inflation). Best investors return is around 14~17%)
        float yield_dividend = 0.03f; // Yield Rate after conquering Financial Freedom (recommend between 0.02 ~0.05)
        float investment_years = 20f;
        String operation = "Salary Share"; // possible values: "Salary Share", "Yield", "Time", "Net Worth"
        

        final byte MONTHS_IN_YEARS = 12;
        final double NEEDED_NETWORTH = salary_retiredYears*MONTHS_IN_YEARS/yield_dividend; // The amount of money the user wishes to have at the end



        // Variables related to the switch/case operations
        double tentative_result = 0;
        double tentative_dividends = 0;
        final byte ITER_MAX = 120;

        switch (operation){
            case "Net Worth": // Calculating final Net Worth based on User Input
                System.out.println("\n\n\nCalculating Final Net Worth...\n");

                CompoundProjector netWorthObj = new CompoundProjector(initial_Investment, salary_productiveYears, monthlySalaryShare, yield_investment, yield_dividend, investment_years);

                double result_end_Investment = netWorthObj.projectNetWorth();

                double monthlyDividends = netWorthObj.projectDividends();

                System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(result_end_Investment));
                System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(monthlyDividends));
                break;

            case "Yield": // Estimating Yield
                System.out.println("\n\n\nCalculating Annual Interest Yield...\n");

                float yieldFloor = 0.0f;
                float yieldCeil = 2f;
                float yieldTentative = 0.05f;
                

                yieldCase: {
                    for(int i=0; i<=ITER_MAX; i++){
                        CompoundProjector yieldOBJ = new CompoundProjector(initial_Investment, salary_productiveYears, monthlySalaryShare, yieldTentative, yield_dividend, investment_years);
                        tentative_result = yieldOBJ.projectNetWorth();
                        tentative_dividends = yieldOBJ.projectDividends();
                        

                        if (NEEDED_NETWORTH * 0.99999 < tentative_result && tentative_result < NEEDED_NETWORTH * 1.00001){ // Almost Equals, with 0.002% Threshold
                            System.out.println("This calculation took " + i + " iterations.");
                            System.out.println("Estimated Yield: " + Math.round(yieldTentative * 10000)/100f + "%");
                            System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(tentative_result));
                            System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(tentative_dividends));
                            break yieldCase;
                        }
                        // Didn't find the correct value, 'binary search' the new yield_investment
                        if (tentative_result < NEEDED_NETWORTH){
                            yieldFloor = yieldTentative;
                            yieldTentative = (yieldFloor + yieldCeil) / 2;
                        }else{
                            yieldCeil = yieldTentative;
                            yieldTentative = (yieldFloor + yieldCeil) / 2;
                        }
                    }

                    // Didn't find a satisfactory answer but we still need to 
                    System.out.println("This calculation would take more than " + ITER_MAX + " iterations.");
                    System.out.println("Estimated Yield: " + Math.round(yieldTentative * 10000)/100f + "%");
                    System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(tentative_result));
                    System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(tentative_dividends));
                }

                break;
            case "Time":
                System.out.println("\n\n\nCalculating Time for Financial Freedom...\n");

                float timeFloor = 0.0f;
                float timeCeil = 250f;
                float timeTentative = 10f;

                timeCase: {
                    for(int i=0; i<=ITER_MAX; i++){
                        CompoundProjector timeOBJ = new CompoundProjector(initial_Investment, salary_productiveYears, monthlySalaryShare, yield_investment, yield_dividend, timeTentative);
                        tentative_result = timeOBJ.projectNetWorth();
                        tentative_dividends = timeOBJ.projectDividends();
                        

                        if (NEEDED_NETWORTH * 0.99999 < tentative_result && tentative_result < NEEDED_NETWORTH * 1.00001){ // Almost Equals, with 0.002% Threshold
                            System.out.println("This calculation took " + i + " iterations.");
                            System.out.println("Estimated Years: " + Math.round(timeTentative * 100)/100f);
                            System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(tentative_result));
                            System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(tentative_dividends));
                            break timeCase;
                        }
                        // Didn't find the correct value, 'binary search' the new timeTentative
                        if (tentative_result < NEEDED_NETWORTH){
                            timeFloor = timeTentative;
                            timeTentative = (timeFloor + timeCeil) / 2;
                        }else{
                            timeCeil = timeTentative;
                            timeTentative = (timeFloor + timeCeil) / 2;
                        }
                    }

                    // Didn't find a satisfactory answer but we still need to 
                    System.out.println("This calculation would take more than " + ITER_MAX + " iterations.");
                    System.out.println("Estimated Years: " + Math.round(timeTentative * 100)/100f);
                    System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(tentative_result));
                    System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(tentative_dividends));
                }
                break;
            case "Salary Share":
                System.out.println("\n\n\nCalculating how much of your Salary you'll need to save...\n");
                
                float shareFloor = 0.0f;
                float shareCeil = 15f;
                float shareTentative = 0.5f;

                shareCase: {
                    for(int i=0; i<=ITER_MAX; i++){
                        CompoundProjector shareOBJ = new CompoundProjector(initial_Investment, salary_productiveYears, shareTentative, yield_investment, yield_dividend, investment_years);
                        tentative_result = shareOBJ.projectNetWorth();
                        tentative_dividends = shareOBJ.projectDividends();                        

                        if (NEEDED_NETWORTH * 0.99999 < tentative_result && tentative_result < NEEDED_NETWORTH * 1.00001){ // Almost Equals, with 0.002% Threshold
                            System.out.println("This calculation took " + i + " iterations.");
                            System.out.println("Estimated Salary Share: " + Math.round(shareTentative * 10000)/100f + "%");
                            System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(tentative_result));
                            System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(tentative_dividends));
                            break shareCase;
                        }
                        // Didn't find the correct value, 'binary search' the new shareTentative
                        if (tentative_result < NEEDED_NETWORTH){
                            shareFloor = shareTentative;
                            shareTentative = (shareFloor + shareCeil) / 2;
                        }else{
                            shareCeil = shareTentative;
                            shareTentative = (shareFloor + shareCeil) / 2;
                        }
                    }                    
                    // Didn't find a satisfactory answer but we still need to 
                    System.out.println("This calculation would take more than " + ITER_MAX + " iterations.");
                    System.out.println("Estimated Salary Share: " + Math.round(shareTentative * 10000)/100f + "%");
                    System.out.println("Final Net Worth: " + NumberFormat.getCurrencyInstance().format(tentative_result));
                    System.out.println("Monthly Dividends: " + NumberFormat.getCurrencyInstance().format(tentative_dividends));
                }
                break;
        }
    }
}
